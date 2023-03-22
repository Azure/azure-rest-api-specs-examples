/** Samples for IotSecuritySolution ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/GetIoTSecuritySolutionsListByIotHubAndRg.json
     */
    /**
     * Sample code: List IoT Security solutions by resource group and IoT Hub.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listIoTSecuritySolutionsByResourceGroupAndIoTHub(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .iotSecuritySolutions()
            .listByResourceGroup(
                "MyRg",
                "properties.iotHubs/any(i eq"
                    + " \"/subscriptions/075423e9-7d33-4166-8bdf-3920b04e3735/resourceGroups/myRg/providers/Microsoft.Devices/IotHubs/FirstIotHub\")",
                com.azure.core.util.Context.NONE);
    }
}
