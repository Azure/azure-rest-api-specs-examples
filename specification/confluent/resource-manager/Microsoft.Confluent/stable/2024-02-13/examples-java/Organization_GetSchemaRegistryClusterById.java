
/**
 * Samples for Organization GetSchemaRegistryClusterById.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/
     * Organization_GetSchemaRegistryClusterById.json
     */
    /**
     * Sample code: Organization_GetSchemaRegistryClusterById.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationGetSchemaRegistryClusterById(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().getSchemaRegistryClusterByIdWithResponse("myResourceGroup", "myOrganization",
            "env-stgcczjp2j3", "lsrc-stgczkq22z", com.azure.core.util.Context.NONE);
    }
}
