
/**
 * Samples for AzureBareMetalStorageInstances ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-
     * preview/examples/AzureBareMetalStorageInstances_ListByResourceGroup.json
     */
    /**
     * Sample code: List all AzureBareMetalStorage instances in a resource group.
     * 
     * @param manager Entry point to BareMetalInfrastructureManager.
     */
    public static void listAllAzureBareMetalStorageInstancesInAResourceGroup(
        com.azure.resourcemanager.baremetalinfrastructure.BareMetalInfrastructureManager manager) {
        manager.azureBareMetalStorageInstances().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
