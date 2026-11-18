
/**
 * Samples for AccountCostControls Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/AccountCostControl/get.json
     */
    /**
     * Sample code: Get an account-owned budget.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getAnAccountOwnedBudget(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accountCostControls().getWithResponse("foundry-resource-group", "foundry-account", "production-agents",
            com.azure.core.util.Context.NONE);
    }
}
