
import com.azure.resourcemanager.kusto.models.PrincipalPermissionsAction;
import com.azure.resourcemanager.kusto.models.Script;
import com.azure.resourcemanager.kusto.models.ScriptLevel;

/**
 * Samples for Scripts Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoScriptsUpdate.json
     */
    /**
     * Sample code: KustoScriptsUpdate.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void kustoScriptsUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        Script resource = manager.scripts().getWithResponse("kustorptest", "kustoCluster", "KustoDatabase8",
            "kustoScript", com.azure.core.util.Context.NONE).getValue();
        resource.update().withScriptUrl("https://mysa.blob.core.windows.net/container/script.txt")
            .withForceUpdateTag("2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe").withContinueOnErrors(true)
            .withScriptLevel(ScriptLevel.DATABASE)
            .withPrincipalPermissionsAction(PrincipalPermissionsAction.REMOVE_PERMISSION_ON_SCRIPT_COMPLETION).apply();
    }
}
