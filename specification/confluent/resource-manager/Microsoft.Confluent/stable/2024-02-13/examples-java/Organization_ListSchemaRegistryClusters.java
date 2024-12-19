
/**
 * Samples for Organization ListSchemaRegistryClusters.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/
     * Organization_ListSchemaRegistryClusters.json
     */
    /**
     * Sample code: Organization_ListSchemaRegistryClusters.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationListSchemaRegistryClusters(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().listSchemaRegistryClusters("myResourceGroup", "myOrganization", "env-stgcczjp2j3", null,
            null, com.azure.core.util.Context.NONE);
    }
}
