import com.azure.resourcemanager.apimanagement.models.TagContract;

/** Samples for Tag Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateTag.json
     */
    /**
     * Sample code: ApiManagementUpdateTag.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        TagContract resource =
            manager
                .tags()
                .getWithResponse("rg1", "apimService1", "temptag", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withDisplayName("temp tag").withIfMatch("*").apply();
    }
}
