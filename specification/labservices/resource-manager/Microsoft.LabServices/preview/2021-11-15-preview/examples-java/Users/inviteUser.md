Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-labservices_1.0.0-beta.2/sdk/labservices/azure-resourcemanager-labservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.labservices.models.InviteBody;

/** Samples for Users Invite. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Users/inviteUser.json
     */
    /**
     * Sample code: inviteUser.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void inviteUser(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager
            .users()
            .invite(
                "testrg123",
                "testlab",
                "testuser",
                new InviteBody().withText("Invitation to lab testlab"),
                Context.NONE);
    }
}
```
