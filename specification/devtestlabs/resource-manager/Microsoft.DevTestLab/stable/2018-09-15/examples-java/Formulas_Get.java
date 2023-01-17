/** Samples for Formulas Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Formulas_Get.json
     */
    /**
     * Sample code: Formulas_Get.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void formulasGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .formulas()
            .getWithResponse("resourceGroupName", "{labName}", "{formulaName}", null, com.azure.core.util.Context.NONE);
    }
}
