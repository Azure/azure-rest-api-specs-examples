import com.azure.resourcemanager.videoanalyzer.models.AccountEncryption;
import com.azure.resourcemanager.videoanalyzer.models.AccountEncryptionKeyType;
import com.azure.resourcemanager.videoanalyzer.models.IotHub;
import com.azure.resourcemanager.videoanalyzer.models.ResourceIdentity;
import com.azure.resourcemanager.videoanalyzer.models.StorageAccount;
import com.azure.resourcemanager.videoanalyzer.models.UserAssignedManagedIdentity;
import com.azure.resourcemanager.videoanalyzer.models.VideoAnalyzerIdentity;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for VideoAnalyzers CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-analyzer-accounts-create-or-update.json
     */
    /**
     * Sample code: Create a Video Analyzer account.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void createAVideoAnalyzerAccount(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager
            .videoAnalyzers()
            .define("contosotv")
            .withRegion("South Central US")
            .withExistingResourceGroup("contoso")
            .withTags(mapOf("tag1", "value1", "tag2", "value2"))
            .withIdentity(
                new VideoAnalyzerIdentity()
                    .withType("UserAssigned")
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1",
                            new UserAssignedManagedIdentity(),
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2",
                            new UserAssignedManagedIdentity(),
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id3",
                            new UserAssignedManagedIdentity())))
            .withStorageAccounts(
                Arrays
                    .asList(
                        new StorageAccount()
                            .withId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Storage/storageAccounts/storage1")
                            .withIdentity(
                                new ResourceIdentity()
                                    .withUserAssignedIdentity(
                                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2"))))
            .withEncryption(new AccountEncryption().withType(AccountEncryptionKeyType.SYSTEM_KEY))
            .withIotHubs(
                Arrays
                    .asList(
                        new IotHub()
                            .withId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Devices/IotHubs/hub1")
                            .withIdentity(
                                new ResourceIdentity()
                                    .withUserAssignedIdentity(
                                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id3")),
                        new IotHub()
                            .withId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Devices/IotHubs/hub2")
                            .withIdentity(
                                new ResourceIdentity()
                                    .withUserAssignedIdentity(
                                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id3"))))
            .create();
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
