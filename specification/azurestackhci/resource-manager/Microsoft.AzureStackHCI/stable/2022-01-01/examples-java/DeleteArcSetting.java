import com.azure.core.util.Context;

/** Samples for ArcSettings Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-01-01/examples/DeleteArcSetting.json
     */
    /**
     * Sample code: Delete ArcSetting.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteArcSetting(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.arcSettings().delete("test-rg", "myCluster", "default", Context.NONE);
    }
}
