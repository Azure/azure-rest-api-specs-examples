
/**
 * Samples for ComplianceResults Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2017-08-01/examples/ComplianceResults/
     * GetComplianceResults_example.json
     */
    /**
     * Sample code: Get compliance results on subscription.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getComplianceResultsOnSubscription(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.complianceResults().getWithResponse("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23",
            "DesignateMoreThanOneOwner", com.azure.core.util.Context.NONE);
    }
}
