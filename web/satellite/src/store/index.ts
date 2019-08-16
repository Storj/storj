// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

import Vue from 'vue';
import Vuex from 'vuex';

import { makeUsersModule } from '@/store/modules/users';
import { makeProjectsModule } from '@/store/modules/projects';
import { projectMembersModule } from '@/store/modules/projectMembers';
import { notificationsModule } from '@/store/modules/notifications';
import { appStateModule } from '@/store/modules/appState';
import { makeApiKeysModule } from '@/store/modules/apiKeys';
import { makeCreditsModule } from '@/store/modules/credits';
import { bucketUsageModule, usageModule } from '@/store/modules/usage';
import { projectPaymentsMethodsModule } from '@/store/modules/paymentMethods';
import { CreditsApiGql } from '@/api/credits';
import { UsersApiGql } from '@/api/users';

Vue.use(Vuex);

export class StoreModule<S> {
    public state: S;
    public mutations: any;
    public actions: any;
    public getters?: any;
}

// TODO: remove it after we will use modules as classes and use some DI framework
const usersApi = new UsersApiGql();
const creditsApi = new CreditsApiGql();

// Satellite store (vuex)
const store = new Vuex.Store({
    modules: {
        usersModule: makeUsersModule(usersApi),
        projectsModule: makeProjectsModule(),
        projectMembersModule,
        notificationsModule,
        appStateModule,
        apiKeysModule: makeApiKeysModule(),
        usageModule,
        bucketUsageModule,
        projectPaymentsMethodsModule,
        creditsModule: makeCreditsModule(creditsApi),
    }
});

export default store;
