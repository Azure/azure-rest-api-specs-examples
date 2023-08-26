/** Samples for Services CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/preview/2023-07-01-preview/examples/Services_CreateOrUpdate.json
     */
    /**
     * Sample code: Services_Create.
     *
     * @param manager Entry point to ApiCenterManager.
     */
    public static void servicesCreate(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager
            .services()
            .define("contoso")
            .withRegion((String) null)
            .withExistingResourceGroup("contoso-resources")
            .create();
    }
}
