
import com.azure.resourcemanager.securityinsights.models.EyesOn;

/**
 * Samples for ProductSettings Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * settings/UpdateEyesOnSetting.json
     */
    /**
     * Sample code: Update EyesOn settings.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        updateEyesOnSettings(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.productSettings().updateWithResponse("myRg", "myWorkspace", "EyesOn",
            new EyesOn().withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\""), com.azure.core.util.Context.NONE);
    }
}
