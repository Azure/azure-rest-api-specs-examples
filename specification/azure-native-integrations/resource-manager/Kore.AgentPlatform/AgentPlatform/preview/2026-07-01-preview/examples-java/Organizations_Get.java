
/**
 * Samples for Organizations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Organizations_Get.json
     */
    /**
     * Sample code: Organizations_Get.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void organizationsGet(com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        manager.organizations().getByResourceGroupWithResponse("contoso-rg", "contoso-org",
            com.azure.core.util.Context.NONE);
    }
}
