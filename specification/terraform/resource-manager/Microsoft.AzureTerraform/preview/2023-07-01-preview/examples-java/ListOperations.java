
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-07-01-preview/ListOperations.json
     */
    /**
     * Sample code: Get a list of operations for a resource provider.
     * 
     * @param manager Entry point to AzureTerraformManager.
     */
    public static void
        getAListOfOperationsForAResourceProvider(com.azure.resourcemanager.terraform.AzureTerraformManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
