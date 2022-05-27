Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.LocalUserInner;
import com.azure.resourcemanager.storage.models.PermissionScope;
import com.azure.resourcemanager.storage.models.SshPublicKey;
import java.util.Arrays;

/** Samples for LocalUsersOperation CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/LocalUserCreate.json
     */
    /**
     * Sample code: CreateLocalUser.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createLocalUser(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getLocalUsersOperations()
            .createOrUpdateWithResponse(
                "res6977",
                "sto2527",
                "user1",
                new LocalUserInner()
                    .withPermissionScopes(
                        Arrays
                            .asList(
                                new PermissionScope()
                                    .withPermissions("rwd")
                                    .withService("file")
                                    .withResourceName("share1"),
                                new PermissionScope()
                                    .withPermissions("rw")
                                    .withService("file")
                                    .withResourceName("share2")))
                    .withHomeDirectory("homedirectory")
                    .withSshAuthorizedKeys(
                        Arrays
                            .asList(new SshPublicKey().withDescription("key name").withKey("ssh-rsa keykeykeykeykey=")))
                    .withHasSshPassword(true),
                Context.NONE);
    }
}
```
