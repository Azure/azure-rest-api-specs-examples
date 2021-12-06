Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.FileShareInner;
import com.azure.resourcemanager.storage.models.AccessPolicy;
import com.azure.resourcemanager.storage.models.SignedIdentifier;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for FileShares Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/FileShareAclsPatch.json
     */
    /**
     * Sample code: UpdateShareAcls.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateShareAcls(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getFileShares()
            .updateWithResponse(
                "res3376",
                "sto328",
                "share6185",
                new FileShareInner()
                    .withSignedIdentifiers(
                        Arrays
                            .asList(
                                new SignedIdentifier()
                                    .withId("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI")
                                    .withAccessPolicy(
                                        new AccessPolicy()
                                            .withStartTime(OffsetDateTime.parse("2021-04-01T08:49:37.0000000Z"))
                                            .withExpiryTime(OffsetDateTime.parse("2021-05-01T08:49:37.0000000Z"))
                                            .withPermission("rwd")))),
                Context.NONE);
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
```
