import com.azure.core.util.Context;

/** Samples for NamedValue Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteNamedValue.json
     */
    /**
     * Sample code: ApiManagementDeleteNamedValue.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteNamedValue(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.namedValues().deleteWithResponse("rg1", "apimService1", "testprop2", "*", Context.NONE);
    }
}
