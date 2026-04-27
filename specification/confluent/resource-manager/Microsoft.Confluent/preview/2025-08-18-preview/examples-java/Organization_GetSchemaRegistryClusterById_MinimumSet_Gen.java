
/**
 * Samples for Organization GetSchemaRegistryClusterById.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Organization_GetSchemaRegistryClusterById_MinimumSet_Gen.json
     */
    /**
     * Sample code: Organization_GetSchemaRegistryClusterById_MinimumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationGetSchemaRegistryClusterByIdMinimumSet(
        com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().getSchemaRegistryClusterByIdWithResponse("rgconfluent", "vcen", "zsbdbdljcfrnxxafcchr",
            "ivjcqxutsnlylxo", com.azure.core.util.Context.NONE);
    }
}
