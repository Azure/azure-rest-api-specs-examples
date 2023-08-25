import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceApplyNetworkConfigurationParameters;

/** Samples for ApiManagementService ApplyNetworkConfigurationUpdates. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementApplyNetworkConfigurationUpdates.json
     */
    /**
     * Sample code: ApiManagementApplyNetworkConfigurationUpdates.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementApplyNetworkConfigurationUpdates(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiManagementServices()
            .applyNetworkConfigurationUpdates(
                "rg1",
                "apimService1",
                new ApiManagementServiceApplyNetworkConfigurationParameters().withLocation("west us"),
                com.azure.core.util.Context.NONE);
    }
}
