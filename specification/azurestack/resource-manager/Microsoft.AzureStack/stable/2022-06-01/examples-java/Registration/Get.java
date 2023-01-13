/** Samples for Registrations GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/Registration/Get.json
     */
    /**
     * Sample code: Returns the properties of an Azure Stack registration.
     *
     * @param manager Entry point to AzureStackManager.
     */
    public static void returnsThePropertiesOfAnAzureStackRegistration(
        com.azure.resourcemanager.azurestack.AzureStackManager manager) {
        manager
            .registrations()
            .getByResourceGroupWithResponse("azurestack", "testregistration", com.azure.core.util.Context.NONE);
    }
}
