
import com.azure.resourcemanager.customerinsights.models.GetImageUploadUrlInput;

/**
 * Samples for Images GetUploadUrlForData.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/
     * ImagesGetUploadUrlForData.json
     */
    /**
     * Sample code: Images_GetUploadUrlForData.
     * 
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void
        imagesGetUploadUrlForData(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.images().getUploadUrlForDataWithResponse(
            "TestHubRG", "sdkTestHub", new GetImageUploadUrlInput().withEntityType("Profile")
                .withEntityTypeName("Contact").withRelativePath("images/profile1.png"),
            com.azure.core.util.Context.NONE);
    }
}
