import com.azure.core.util.Context;

/** Samples for Certificate ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListCertificates.json
     */
    /**
     * Sample code: ApiManagementListCertificates.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListCertificates(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.certificates().listByService("rg1", "apimService1", null, null, null, null, Context.NONE);
    }
}
