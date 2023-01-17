/** Samples for ArmTemplates List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArmTemplates_List.json
     */
    /**
     * Sample code: ArmTemplates_List.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void armTemplatesList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .armTemplates()
            .list(
                "resourceGroupName",
                "{labName}",
                "{artifactSourceName}",
                null,
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
