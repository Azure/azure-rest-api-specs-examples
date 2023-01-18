/** Samples for Accounts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_Delete.json
     */
    /**
     * Sample code: Accounts_Delete.
     *
     * @param manager Entry point to PurviewManager.
     */
    public static void accountsDelete(com.azure.resourcemanager.purview.PurviewManager manager) {
        manager.accounts().delete("SampleResourceGroup", "account1", com.azure.core.util.Context.NONE);
    }
}
