
import com.azure.resourcemanager.securityinsights.models.IoTCheckRequirements;

/**
 * Samples for DataConnectorsCheckRequirementsOperation Post.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * dataConnectors/CheckRequirementsIoT.json
     */
    /**
     * Sample code: Check requirements for IoT.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        checkRequirementsForIoT(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectorsCheckRequirementsOperations().postWithResponse("myRg", "myWorkspace",
            new IoTCheckRequirements().withSubscriptionId("c0688291-89d7-4bed-87a2-a7b1bff43f4c"),
            com.azure.core.util.Context.NONE);
    }
}
