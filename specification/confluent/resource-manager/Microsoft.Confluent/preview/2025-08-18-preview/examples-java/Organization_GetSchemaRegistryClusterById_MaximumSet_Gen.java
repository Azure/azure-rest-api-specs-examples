
/**
 * Samples for Organization GetSchemaRegistryClusterById.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Organization_GetSchemaRegistryClusterById_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organization_GetSchemaRegistryClusterById_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationGetSchemaRegistryClusterByIdMaximumSet(
        com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().getSchemaRegistryClusterByIdWithResponse("rgconfluent", "hmhbrtw",
            "ztozszmpzhwevkpmaxslloijkicwt", "stfqijternpuzpleowkrbgzuutsgp", com.azure.core.util.Context.NONE);
    }
}
