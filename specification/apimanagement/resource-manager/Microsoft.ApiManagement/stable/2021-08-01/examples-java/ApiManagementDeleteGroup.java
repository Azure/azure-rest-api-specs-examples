import com.azure.core.util.Context;

/** Samples for Group Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteGroup.json
     */
    /**
     * Sample code: ApiManagementDeleteGroup.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteGroup(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.groups().deleteWithResponse("rg1", "apimService1", "aadGroup", "*", Context.NONE);
    }
}
