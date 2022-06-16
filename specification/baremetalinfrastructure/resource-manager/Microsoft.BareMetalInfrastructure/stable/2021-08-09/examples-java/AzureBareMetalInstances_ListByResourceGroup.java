import com.azure.core.util.Context;

/** Samples for AzureBareMetalInstances ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/stable/2021-08-09/examples/AzureBareMetalInstances_ListByResourceGroup.json
     */
    /**
     * Sample code: List all AzureBareMetal instances in a resource group.
     *
     * @param manager Entry point to BareMetalInfrastructureManager.
     */
    public static void listAllAzureBareMetalInstancesInAResourceGroup(
        com.azure.resourcemanager.baremetalinfrastructure.BareMetalInfrastructureManager manager) {
        manager.azureBareMetalInstances().listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
