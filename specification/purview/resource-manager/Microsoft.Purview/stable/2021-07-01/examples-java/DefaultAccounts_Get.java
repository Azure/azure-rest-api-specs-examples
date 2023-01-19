import com.azure.resourcemanager.purview.models.ScopeType;
import java.util.UUID;

/** Samples for DefaultAccounts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/DefaultAccounts_Get.json
     */
    /**
     * Sample code: DefaultAccounts_Get.
     *
     * @param manager Entry point to PurviewManager.
     */
    public static void defaultAccountsGet(com.azure.resourcemanager.purview.PurviewManager manager) {
        manager
            .defaultAccounts()
            .getWithResponse(
                UUID.fromString("11733A4E-BA84-46FF-91D1-AFF1A3215A90"),
                ScopeType.TENANT,
                "11733A4E-BA84-46FF-91D1-AFF1A3215A90",
                com.azure.core.util.Context.NONE);
    }
}
