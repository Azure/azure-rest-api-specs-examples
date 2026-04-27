
/**
 * Samples for ServiceGroupMemberRelationships Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-09-01-preview/ServiceGroupMemberRelationships_Delete.json
     */
    /**
     * Sample code: ServiceGroupMemberRelationships_Delete.
     * 
     * @param manager Entry point to RelationshipsManager.
     */
    public static void
        serviceGroupMemberRelationshipsDelete(com.azure.resourcemanager.relationships.RelationshipsManager manager) {
        manager.serviceGroupMemberRelationships().delete(
            "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.DocumentDb/databaseAccounts/test-db-account",
            "sg1", com.azure.core.util.Context.NONE);
    }
}
