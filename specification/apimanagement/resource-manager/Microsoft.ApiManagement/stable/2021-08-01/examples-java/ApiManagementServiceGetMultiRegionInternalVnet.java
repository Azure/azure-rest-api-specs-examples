import com.azure.core.util.Context;

/** Samples for ApiManagementService GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementServiceGetMultiRegionInternalVnet.json
     */
    /**
     * Sample code: ApiManagementServiceGetMultiRegionInternalVnet.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementServiceGetMultiRegionInternalVnet(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiManagementServices().getByResourceGroupWithResponse("rg1", "apimService1", Context.NONE);
    }
}
