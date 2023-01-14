import com.azure.resourcemanager.databox.models.IdentityProperties;
import com.azure.resourcemanager.databox.models.JobResource;
import com.azure.resourcemanager.databox.models.KekType;
import com.azure.resourcemanager.databox.models.KeyEncryptionKey;
import com.azure.resourcemanager.databox.models.ResourceIdentity;
import com.azure.resourcemanager.databox.models.UpdateJobDetails;
import com.azure.resourcemanager.databox.models.UserAssignedIdentity;
import com.azure.resourcemanager.databox.models.UserAssignedProperties;
import java.util.HashMap;
import java.util.Map;

/** Samples for Jobs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2021-03-01/examples/JobsPatchSystemAssignedToUserAssigned.json
     */
    /**
     * Sample code: JobsPatchSystemAssignedToUserAssigned.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsPatchSystemAssignedToUserAssigned(com.azure.resourcemanager.databox.DataBoxManager manager) {
        JobResource resource =
            manager
                .jobs()
                .getByResourceGroupWithResponse("SdkRg9765", "SdkJob2965", null, com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withIdentity(
                new ResourceIdentity()
                    .withType("SystemAssigned,UserAssigned")
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.ManagedIdentity/userAssignedIdentities/sdkIdentity",
                            new UserAssignedIdentity())))
            .withDetails(
                new UpdateJobDetails()
                    .withKeyEncryptionKey(
                        new KeyEncryptionKey()
                            .withKekType(KekType.CUSTOMER_MANAGED)
                            .withIdentityProperties(
                                new IdentityProperties()
                                    .withType("UserAssigned")
                                    .withUserAssigned(
                                        new UserAssignedProperties()
                                            .withResourceId(
                                                "/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.ManagedIdentity/userAssignedIdentities/sdkIdentity")))
                            .withKekUrl("https://sdkkeyvault.vault.azure.net/keys/SSDKEY/")
                            .withKekVaultResourceId(
                                "/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.KeyVault/vaults/SDKKeyVault")))
            .apply();
    }

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
