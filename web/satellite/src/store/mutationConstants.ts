// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

export const USER_MUTATIONS = {
    SET_USER: 'setUser',
    UPDATE_USER: 'updateUser',
    CLEAR: 'clearUser',
};

export const PROJECTS_MUTATIONS = {
    CREATE: 'CREATE_PROJECT',
    DELETE: 'DELETE_PROJECT',
    UPDATE: 'UPDATE_PROJECT',
    FETCH: 'FETCH_PROJECTS',
    SELECT: 'SELECT_PROJECT',
    CLEAR: 'CLEAR_PROJECTS',
};

export const PROJECT_MEMBER_MUTATIONS = {
    FETCH: 'FETCH_MEMBERS',
    TOGGLE_SELECTION: 'TOGGLE_SELECTION',
    CLEAR_SELECTION: 'CLEAR_SELECTION',
    DELETE: 'DELETE_MEMBERS',
    CLEAR: 'CLEAR_MEMBERS',
    CHANGE_SORT_ORDER: 'CHANGE_SORT_ORDER',
    CHANGE_SORT_ORDER_DIRECTION: 'CHANGE_SORT_ORDER_DIRECTION',
    SET_SEARCH_QUERY: 'SET_SEARCH_QUERY',
    SET_PAGE: 'SET_PROJECT_MEMBERS_PAGE',
};

export const API_KEYS_MUTATIONS = {
    FETCH: 'FETCH_API_KEYS',
    ADD: 'ADD_API_KEY',
    DELETE: 'DELETE_API_KEY',
    TOGGLE_SELECTION: 'TOGGLE_SELECTION',
    CLEAR_SELECTION: 'CLEAR_SELECTION',
    CLEAR: 'CLEAR_API_KEYS',
};

export const PROJECT_USAGE_MUTATIONS = {
    FETCH: 'FETCH_PROJECT_USAGE',
    SET_DATE: 'SET_DATE_PROJECT_USAGE',
    CLEAR: 'CLEAR_PROJECT_USAGE'
};

export const BUCKET_USAGE_MUTATIONS = {
    FETCH: 'FETCH_BUCKET_USAGES',
    SET_SEARCH: 'SET_SEARCH_BUCKET_USAGE',
    SET_PAGE: 'SET_PAGE_BUCKET_USAGE',
    CLEAR: 'CLEAR_BUCKET_USAGES'
};

export const CREDIT_USAGE_MUTATIONS = {
    FETCH: 'FETCH_CREDIT_USAGE',
};

export const NOTIFICATION_MUTATIONS = {
    ADD: 'ADD_NOTIFICATION',
    DELETE: 'DELETE_NOTIFICATION',
    PAUSE: 'PAUSE_NOTIFICATION',
    RESUME: 'RESUME_NOTIFICATION',
    CLEAR: 'CLEAR_NOTIFICATIONS',
};

export const APP_STATE_MUTATIONS = {
    TOGGLE_ADD_TEAMMEMBER_POPUP: 'TOGGLE_ADD_TEAMMEMBER_POPUP',
    TOGGLE_NEW_PROJECT_POPUP: 'TOGGLE_NEW_PROJECT_POPUP',
    TOGGLE_PROJECT_DROPDOWN: 'TOGGLE_PROJECT_DROPDOWN',
    TOGGLE_ACCOUNT_DROPDOWN: 'TOGGLE_ACCOUNT_DROPDOWN',
    TOGGLE_DELETE_PROJECT_DROPDOWN: 'TOGGLE_DELETE_PROJECT_DROPDOWN',
    TOGGLE_DELETE_ACCOUNT_DROPDOWN: 'TOGGLE_DELETE_ACCOUNT_DROPDOWN',
    TOGGLE_SORT_PM_BY_DROPDOWN: 'TOGGLE_SORT_PM_BY_DROPDOWN',
    TOGGLE_NEW_API_KEY_POPUP: 'TOGGLE_NEW_API_KEY_POPUP',
    TOGGLE_SUCCESSFUL_REGISTRATION_POPUP: 'TOGGLE_SUCCESSFUL_REGISTRATION_POPUP',
    TOGGLE_SUCCESSFUL_PROJECT_CREATION_POPUP: 'TOGGLE_SUCCESSFUL_PROJECT_CREATION_POPUP',
    TOGGLE_EDIT_PROFILE_POPUP: 'TOGGLE_EDIT_PROFILE_POPUP',
    TOGGLE_CHANGE_PASSWORD_POPUP: 'TOGGLE_CHANGE_PASSWORD_POPUP',
    SHOW_DELETE_PAYMENT_METHOD_POPUP: 'SHOW_DELETE_PAYMENT_METHOD_POPUP',
    SHOW_SET_DEFAULT_PAYMENT_METHOD_POPUP: 'SHOW_SET_DEFAULT_PAYMENT_METHOD_POPUP',
    CLOSE_ALL: 'CLOSE_ALL',
    CHANGE_STATE: 'CHANGE_STATE',
};

export const PROJECT_PAYMENT_METHODS_MUTATIONS = {
    FETCH: 'FETCH',
    CLEAR: 'CLEAR',
};
