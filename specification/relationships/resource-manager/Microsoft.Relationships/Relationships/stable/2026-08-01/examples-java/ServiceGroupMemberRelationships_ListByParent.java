
/**
 * Samples for ServiceGroupMemberRelationships ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01/ServiceGroupMemberRelationships_ListByParent.json
     */
    /**
     * Sample code: ServiceGroupMemberRelationships_ListByParent.
     * 
     * @param manager Entry point to RelationshipsManager.
     */
    public static void serviceGroupMemberRelationshipsListByParent(
        com.azure.resourcemanager.relationships.RelationshipsManager manager) {
        manager.serviceGroupMemberRelationships().listByParent(
            "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.DocumentDb/databaseAccounts/test-db-account",
            com.azure.core.util.Context.NONE);
    }
}
