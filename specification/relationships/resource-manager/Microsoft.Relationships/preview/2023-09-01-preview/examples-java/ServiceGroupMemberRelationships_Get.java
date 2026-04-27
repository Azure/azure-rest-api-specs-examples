
/**
 * Samples for ServiceGroupMemberRelationships Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-09-01-preview/ServiceGroupMemberRelationships_Get.json
     */
    /**
     * Sample code: ServiceGroupMemberRelationships_Get.
     * 
     * @param manager Entry point to RelationshipsManager.
     */
    public static void
        serviceGroupMemberRelationshipsGet(com.azure.resourcemanager.relationships.RelationshipsManager manager) {
        manager.serviceGroupMemberRelationships().getWithResponse(
            "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.DocumentDb/databaseAccounts/test-db-account",
            "sg1", com.azure.core.util.Context.NONE);
    }
}
