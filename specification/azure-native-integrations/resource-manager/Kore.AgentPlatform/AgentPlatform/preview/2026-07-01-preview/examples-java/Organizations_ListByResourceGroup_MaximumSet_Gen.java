
/**
 * Samples for Organizations ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Organizations_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void organizationsListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        manager.organizations().listByResourceGroup("contoso-rg", com.azure.core.util.Context.NONE);
    }
}
