// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

import Vuex from 'vuex';

import { RouteConfig, router } from '@/router';
import { makeApiKeysModule } from '@/store/modules/apiKeys';
import { appStateModule } from '@/store/modules/appState';
import { makeBucketsModule } from '@/store/modules/buckets';
import { makeNotificationsModule } from '@/store/modules/notifications';
import { makePaymentsModule } from '@/store/modules/payments';
import { makeProjectMembersModule } from '@/store/modules/projectMembers';
import { makeProjectsModule } from '@/store/modules/projects';
import { makeUsageModule } from '@/store/modules/usage';
import { makeUsersModule } from '@/store/modules/users';
import { APP_STATE_MUTATIONS } from '@/store/mutationConstants';
import { User } from '@/types/users';
import { APP_STATE_ACTIONS } from '@/utils/constants/actionNames';
import { AppState } from '@/utils/constants/appStateEnum';
import { NotificatorPlugin } from '@/utils/plugins/notificator';
import { SegmentioPlugin } from '@/utils/plugins/segment';
import DashboardArea from '@/views/DashboardArea.vue';
import { createLocalVue, shallowMount } from '@vue/test-utils';

import { ApiKeysMock } from '../mock/api/apiKeys';
import { BucketsMock } from '../mock/api/buckets';
import { PaymentsMock } from '../mock/api/payments';
import { ProjectMembersApiMock } from '../mock/api/projectMembers';
import { ProjectsApiMock } from '../mock/api/projects';
import { ProjectUsageMock } from '../mock/api/usage';
import { UsersApiMock } from '../mock/api/users';

const localVue = createLocalVue();
const segmentioPlugin = new SegmentioPlugin();
const notificationPlugin = new NotificatorPlugin();
localVue.use(Vuex);
localVue.use(segmentioPlugin);
localVue.use(notificationPlugin);

const usersApi = new UsersApiMock();
const projectsApi = new ProjectsApiMock();

usersApi.setMockUser(new User('1', '2', '3', '4', '5'));
projectsApi.setMockProjects([]);

const usersModule = makeUsersModule(usersApi);
const projectsModule = makeProjectsModule(projectsApi);
const apiKeysModule = makeApiKeysModule(new ApiKeysMock());
const teamMembersModule = makeProjectMembersModule(new ProjectMembersApiMock());
const bucketsModule = makeBucketsModule(new BucketsMock());
const usageModule = makeUsageModule(new ProjectUsageMock());
const notificationsModule = makeNotificationsModule();
const paymentsModule = makePaymentsModule(new PaymentsMock());

const store = new Vuex.Store({
    modules: {
        notificationsModule,
        usageModule,
        bucketsModule,
        apiKeysModule,
        usersModule,
        projectsModule,
        appStateModule,
        teamMembersModule,
        paymentsModule,
    },
});

describe('Dashboard', () => {
    beforeEach(() => {
        jest.resetAllMocks();
    });

    it('renders correctly when data is loading', () => {
        const wrapper = shallowMount(DashboardArea, {
            store,
            localVue,
            router,
        });

        expect(wrapper).toMatchSnapshot();
        expect(wrapper.findAll('.loading-overlay.active').length).toBe(1);
        expect(wrapper.findAll('.dashboard-container__wrap').length).toBe(0);
    });

    it('renders correctly without project and with payment method', async () => {
        store.commit(APP_STATE_MUTATIONS.SHOW_CONTENT_BLUR);

        const wrapper = shallowMount(DashboardArea, {
            store,
            localVue,
            router,
        });

        expect(wrapper).toMatchSnapshot();

        await wrapper.find('.dashboard-container__blur-area__button').trigger('click');

        expect(wrapper).toMatchSnapshot();
    });

    it('renders correctly when data is loaded', () => {
        store.dispatch(APP_STATE_ACTIONS.CHANGE_STATE, AppState.LOADED);

        const wrapper = shallowMount(DashboardArea, {
            store,
            localVue,
            router,
        });

        expect(wrapper).toMatchSnapshot();
        expect(wrapper.findAll('.loading-overlay active').length).toBe(0);
        expect(wrapper.findAll('.dashboard-container__wrap').length).toBe(1);
    });

    it('loads routes correctly when authorithed without project with available routes', async () => {
        const availableWithoutProject = [
            RouteConfig.Account.with(RouteConfig.Billing).path,
            RouteConfig.Account.with(RouteConfig.Settings).path,
        ];

        for (let i = 0; i < availableWithoutProject.length; i++) {
            const wrapper = await shallowMount(DashboardArea, {
                localVue,
                router,
                store,
            });

            setTimeout(() => {
                expect(wrapper.vm.$router.currentRoute.path).toBe(availableWithoutProject[i]);
            }, 50);
        }
    });

    it('loads routes correctly when authorithed without project with unavailable routes', async () => {
        const unavailableWithoutProject = [
            RouteConfig.ApiKeys.path,
            RouteConfig.Buckets.path,
            RouteConfig.Team.path,
            RouteConfig.ProjectDashboard.with(RouteConfig.UsageReport).path,
        ];

        for (let i = 0; i < unavailableWithoutProject.length; i++) {
            await router.push(unavailableWithoutProject[i]);

            const wrapper = await shallowMount(DashboardArea, {
                localVue,
                router,
                store,
            });

            setTimeout(() => {
                expect(wrapper.vm.$router.currentRoute.path).toBe(RouteConfig.ProjectDashboard.with(RouteConfig.ProjectDetails).path);
            }, 50);
        }

    });

    it('loads routes correctly when not authorithed', () => {
        const wrapper = shallowMount(DashboardArea, {
            store,
            localVue,
            router,
        });

        setTimeout(() => {
            expect(wrapper.vm.$router.currentRoute.path).toBe(RouteConfig.Login.path);
        }, 50);
    });
});
