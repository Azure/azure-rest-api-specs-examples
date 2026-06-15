
/**
 * Samples for Formulas Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Formulas_Delete.json
     */
    /**
     * Sample code: Formulas_Delete.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void formulasDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.formulas().deleteWithResponse("resourceGroupName", "{labName}", "{formulaName}",
            com.azure.core.util.Context.NONE);
    }
}
