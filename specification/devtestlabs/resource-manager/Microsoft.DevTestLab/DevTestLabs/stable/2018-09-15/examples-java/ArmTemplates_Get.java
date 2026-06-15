
/**
 * Samples for ArmTemplates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/ArmTemplates_Get.json
     */
    /**
     * Sample code: ArmTemplates_Get.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void armTemplatesGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.armTemplates().getWithResponse("resourceGroupName", "{labName}", "{artifactSourceName}",
            "{armTemplateName}", null, com.azure.core.util.Context.NONE);
    }
}
