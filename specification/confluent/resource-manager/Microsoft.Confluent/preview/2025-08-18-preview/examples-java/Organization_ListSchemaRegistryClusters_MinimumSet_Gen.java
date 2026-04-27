
/**
 * Samples for Organization ListSchemaRegistryClusters.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Organization_ListSchemaRegistryClusters_MinimumSet_Gen.json
     */
    /**
     * Sample code: Organization_ListSchemaRegistryClusters_MinimumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationListSchemaRegistryClustersMinimumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().listSchemaRegistryClusters("rgconfluent", "npek", "tdtxr", null, null,
            com.azure.core.util.Context.NONE);
    }
}
