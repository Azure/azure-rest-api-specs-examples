import com.azure.resourcemanager.apimanagement.models.VersioningScheme;

/** Samples for ApiVersionSet CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiVersionSet.json
     */
    /**
     * Sample code: ApiManagementCreateApiVersionSet.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiVersionSet(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiVersionSets()
            .define("api1")
            .withExistingService("rg1", "apimService1")
            .withDisplayName("api set 1")
            .withVersioningScheme(VersioningScheme.SEGMENT)
            .withDescription("Version configuration")
            .create();
    }
}
