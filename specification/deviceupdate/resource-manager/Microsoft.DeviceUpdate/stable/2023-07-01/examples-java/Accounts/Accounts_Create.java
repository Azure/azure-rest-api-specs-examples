
import com.azure.resourcemanager.deviceupdate.models.Encryption;
import com.azure.resourcemanager.deviceupdate.models.ManagedServiceIdentity;
import com.azure.resourcemanager.deviceupdate.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.deviceupdate.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Accounts Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/Accounts/
     * Accounts_Create.json
     */
    /**
     * Sample code: Creates or updates Account.
     * 
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void createsOrUpdatesAccount(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.accounts().define("contoso").withRegion("westus2").withExistingResourceGroup("test-rg")
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1",
                    new UserAssignedIdentity())))
            .withEncryption(new Encryption().withKeyVaultKeyUri("fakeTokenPlaceholder").withUserAssignedIdentity(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1"))
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
