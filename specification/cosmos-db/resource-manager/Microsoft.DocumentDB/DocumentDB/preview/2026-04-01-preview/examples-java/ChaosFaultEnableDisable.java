
import com.azure.resourcemanager.cosmos.fluent.models.ChaosFaultResourceInner;
import com.azure.resourcemanager.cosmos.models.SupportedActions;

/**
 * Samples for ChaosFault EnableDisable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/ChaosFaultEnableDisable.json
     */
    /**
     * Sample code: ChaosFaultEnableDisable.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void chaosFaultEnableDisable(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getChaosFaults().enableDisable("myResourceGroupName", "myAccountName",
            "ServiceUnavailability",
            new ChaosFaultResourceInner().withAction(SupportedActions.ENABLE).withRegion("EastUS")
                .withDatabaseName("testDatabase").withContainerName("testCollection"),
            com.azure.core.util.Context.NONE);
    }
}
