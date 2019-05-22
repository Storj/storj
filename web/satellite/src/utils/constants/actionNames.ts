// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

export const APP_STATE_ACTIONS = {
    TOGGLE_TEAM_MEMBERS: 'toggleAddTeamMembersPopup',
    TOGGLE_NEW_PROJ : 'toggleNewProjectPopup',
    TOGGLE_PROJECTS: 'toggleProjectsDropdown',
    TOGGLE_ACCOUNT: 'toggleAccountDropdown',
    TOGGLE_DEL_PROJ: 'toggleDeleteProjectPopup',
    TOGGLE_DEL_ACCOUNT: 'toggleDeleteAccountPopup',
    TOGGLE_NEW_API_KEY: 'toggleNewAPIKeyPopup',
    TOGGLE_SORT_PM_BY_DROPDOWN: 'toggleSortProjectMembersByPopup',
    TOGGLE_SUCCESSFUL_REGISTRATION_POPUP: 'toggleSuccessfulRegistrationPopup',
    TOGGLE_SUCCESSFUL_PROJECT_CREATION_POPUP: 'toggleSuccessfulProjectCreationPopup',
    TOGGLE_EDIT_PROFILE_POPUP: 'toggleEditProfilePopup',
    TOGGLE_CHANGE_PASSWORD_POPUP: 'toggleChangePasswordPopup',
    CLOSE_POPUPS: 'closePopups',
    CHANGE_STATE: 'changeFetchState',
};

export const NOTIFICATION_ACTIONS = {
    SUCCESS: 'success',
    ERROR: 'error',
    NOTIFY: 'notify',
    DELETE: 'deleteNotification',
    PAUSE: 'pauseNotification',
    RESUME: 'resumeNotification',
    CLEAR: 'clearNotifications',
};

export const PM_ACTIONS = {
    ADD: 'addProjectMembers',
    DELETE: 'deleteProjectMembers',
    TOGGLE_SELECTION: 'toggleProjectMemberSelection',
    CLEAR_SELECTION: 'clearProjectMemberSelection',
    FETCH: 'fetchProjectMembers',
    CLEAR: 'clearProjectMembers',
    SET_SEARCH_QUERY: 'setProjectMembersSearchQuery',
    SET_SORT_BY: 'setProjectMembersSortingBy',
    CLEAR_OFFSET: 'clearProjectMembersOffset'
};

export const PROJETS_ACTIONS = {
    FETCH: 'fetchProjects',
    CREATE: 'createProject',
    SELECT: 'selectProject',
    UPDATE: 'updateProject',
    DELETE: 'deleteProject',
    CLEAR: 'clearProjects',
};

export const USER_ACTIONS = {
    UPDATE: 'updateAccount',
    CHANGE_PASSWORD: 'changePassword',
    DELETE: 'deleteAccount',
    GET: 'getUser',
    CLEAR: 'clearUser',
    ACTIVATE: 'activateAccount',
};

export const API_KEYS_ACTIONS = {
    FETCH: 'fetchAPIKeys',
    CREATE: 'createAPIKey',
    DELETE: 'deleteAPIKey',
    CLEAR: 'clearAPIKeys',
    TOGGLE_SELECTION: 'toggleAPIKeySelection',
    CLEAR_SELECTION: 'clearAPIKeySelection'
};

export const PROJECT_USAGE_ACTIONS = {
    FETCH: 'fetchProjectUsage',
    FETCH_CURRENT_ROLLUP: 'fetchCurrentProjectUsage',
    FETCH_PREVIOUS_ROLLUP: 'fetchPreviousProjectUsage',
    CLEAR: 'clearProjectUsage',
};

export const BUCKET_USAGE_ACTIONS = {
    FETCH: 'fetchBucketUsages',
    SET_SEARCH: 'setSearchBucketUsage',
    CLEAR: 'clearBucketUsages'
};

export const PROJECT_INVOICE_ACTIONS = {
    FETCH: 'fetchProjectInvoices',
    CLEAR: 'clearProjectInvoices'
};
