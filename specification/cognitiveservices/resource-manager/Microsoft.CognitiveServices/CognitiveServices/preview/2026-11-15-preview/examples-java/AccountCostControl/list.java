
/**
 * Samples for AccountCostControls List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/AccountCostControl/list.json
     */
    /**
     * Sample code: List budgets owned by an account.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listBudgetsOwnedByAnAccount(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accountCostControls().list("foundry-resource-group", "foundry-account",
            com.azure.core.util.Context.NONE);
    }
}
