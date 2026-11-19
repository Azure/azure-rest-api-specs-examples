
/**
 * Samples for AccountCostControls CheckExistence.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/AccountCostControl/checkExistence.json
     */
    /**
     * Sample code: Check for a budget without returning its body.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void checkForABudgetWithoutReturningItsBody(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accountCostControls().checkExistenceWithResponse("foundry-resource-group", "foundry-account",
            "production-agents", com.azure.core.util.Context.NONE);
    }
}
