/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntPredicament,
    EntPredicamentFromJSON,
    EntPredicamentFromJSONTyped,
    EntPredicamentToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntStaffEdges
 */
export interface EntStaffEdges {
    /**
     * StaffPredicament holds the value of the staff_predicament edge.
     * @type {Array<EntPredicament>}
     * @memberof EntStaffEdges
     */
    staffPredicament?: Array<EntPredicament>;
}

export function EntStaffEdgesFromJSON(json: any): EntStaffEdges {
    return EntStaffEdgesFromJSONTyped(json, false);
}

export function EntStaffEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStaffEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'staffPredicament': !exists(json, 'staffPredicament') ? undefined : ((json['staffPredicament'] as Array<any>).map(EntPredicamentFromJSON)),
    };
}

export function EntStaffEdgesToJSON(value?: EntStaffEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'staffPredicament': value.staffPredicament === undefined ? undefined : ((value.staffPredicament as Array<any>).map(EntPredicamentToJSON)),
    };
}

