import com.azure.resourcemanager.purview.models.ScopeType;
import java.util.UUID;

/** Samples for DefaultAccounts Remove. */
public final class Main {
    /*
     * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/DefaultAccounts_Remove.json
     */
    /**
     * Sample code: DefaultAccounts_Remove.
     *
     * @param manager Entry point to PurviewManager.
     */
    public static void defaultAccountsRemove(com.azure.resourcemanager.purview.PurviewManager manager) {
        manager
            .defaultAccounts()
            .removeWithResponse(
                UUID.fromString("11733A4E-BA84-46FF-91D1-AFF1A3215A90"),
                ScopeType.TENANT,
                "11733A4E-BA84-46FF-91D1-AFF1A3215A90",
                com.azure.core.util.Context.NONE);
    }
}
