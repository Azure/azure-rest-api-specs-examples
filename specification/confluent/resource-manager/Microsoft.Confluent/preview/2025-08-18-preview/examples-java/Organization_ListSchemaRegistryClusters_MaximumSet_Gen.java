
/**
 * Samples for Organization ListSchemaRegistryClusters.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Organization_ListSchemaRegistryClusters_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organization_ListSchemaRegistryClusters_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationListSchemaRegistryClustersMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().listSchemaRegistryClusters("rgconfluent", "vkzifcygqhoewuixdmmg", "psxriyxxbjnctgeohah",
            3, "npqeazvityguunrpgbumrqivvq", com.azure.core.util.Context.NONE);
    }
}
