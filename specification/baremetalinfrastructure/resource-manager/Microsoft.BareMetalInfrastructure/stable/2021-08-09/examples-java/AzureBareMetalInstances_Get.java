import com.azure.core.util.Context;

/** Samples for AzureBareMetalInstances GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/stable/2021-08-09/examples/AzureBareMetalInstances_Get.json
     */
    /**
     * Sample code: Get an AzureBareMetal instance.
     *
     * @param manager Entry point to BareMetalInfrastructureManager.
     */
    public static void getAnAzureBareMetalInstance(
        com.azure.resourcemanager.baremetalinfrastructure.BareMetalInfrastructureManager manager) {
        manager
            .azureBareMetalInstances()
            .getByResourceGroupWithResponse("myResourceGroup", "myAzureBareMetalInstance", Context.NONE);
    }
}
