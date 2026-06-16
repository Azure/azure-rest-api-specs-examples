
/**
 * Samples for Formulas List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Formulas_List.json
     */
    /**
     * Sample code: Formulas_List.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void formulasList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.formulas().list("resourceGroupName", "{labName}", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
