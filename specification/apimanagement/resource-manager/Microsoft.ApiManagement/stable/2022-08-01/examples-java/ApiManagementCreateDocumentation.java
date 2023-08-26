/** Samples for Documentation CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateDocumentation.json
     */
    /**
     * Sample code: ApiManagementCreateDocumentation.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateDocumentation(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .documentations()
            .define("57d1f7558aa04f15146d9d8a")
            .withExistingService("rg1", "apimService1")
            .withTitle("Title")
            .withContent("content")
            .create();
    }
}
