
/**
 * Samples for Organizations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Organizations_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void
        organizationsGetMaximumSetGen(com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        manager.organizations().getByResourceGroupWithResponse("contoso-rg", "contoso-org",
            com.azure.core.util.Context.NONE);
    }
}
