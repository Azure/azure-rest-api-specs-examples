/** Samples for Certificate CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateCertificate.json
     */
    /**
     * Sample code: ApiManagementCreateCertificate.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateCertificate(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .certificates()
            .define("tempcert")
            .withExistingService("rg1", "apimService1")
            .withData("****************Base 64 Encoded Certificate *******************************")
            .withPassword("****Certificate Password******")
            .create();
    }
}
