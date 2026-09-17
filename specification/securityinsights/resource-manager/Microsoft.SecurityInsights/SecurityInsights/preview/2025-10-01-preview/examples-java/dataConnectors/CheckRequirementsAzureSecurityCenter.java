
import com.azure.resourcemanager.securityinsights.models.ASCCheckRequirements;

/**
 * Samples for DataConnectorsCheckRequirementsOperation Post.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/CheckRequirementsAzureSecurityCenter.json
     */
    /**
     * Sample code: Check requirements for ASC.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        checkRequirementsForASC(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectorsCheckRequirementsOperations().postWithResponse("myRg", "myWorkspace",
            new ASCCheckRequirements().withSubscriptionId("c0688291-89d7-4bed-87a2-a7b1bff43f4c"),
            com.azure.core.util.Context.NONE);
    }
}
