
import com.azure.resourcemanager.relationships.models.ServiceGroupMemberRelationshipPropertiesV2;

/**
 * Samples for ServiceGroupMemberRelationships CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01/ServiceGroupMemberRelationships_CreateOrUpdate.json
     */
    /**
     * Sample code: ServiceGroupMemberRelationships_CreateOrUpdate.
     * 
     * @param manager Entry point to RelationshipsManager.
     */
    public static void serviceGroupMemberRelationshipsCreateOrUpdate(
        com.azure.resourcemanager.relationships.RelationshipsManager manager) {
        manager.serviceGroupMemberRelationships().define("sg1").withExistingResourceUri(
            "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.DocumentDb/databaseAccounts/test-db-account")
            .withProperties(new ServiceGroupMemberRelationshipPropertiesV2()
                .withSourceId("/providers/Microsoft.Management/serviceGroups/sg1")
                .withSourceTenant("72f988bf-86f1-41af-91ab-2d7cd011db47"))
            .create();
    }
}
