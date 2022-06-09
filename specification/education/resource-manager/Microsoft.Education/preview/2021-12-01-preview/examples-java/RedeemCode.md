```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.education.models.RedeemRequest;

/** Samples for ResourceProvider RedeemInvitationCode. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/RedeemCode.json
     */
    /**
     * Sample code: RedeemCode.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void redeemCode(com.azure.resourcemanager.education.EducationManager manager) {
        manager
            .resourceProviders()
            .redeemInvitationCodeWithResponse(
                new RedeemRequest().withRedeemCode("exampleRedeemCode").withFirstName("test").withLastName("user"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-education_1.0.0-beta.1/sdk/education/azure-resourcemanager-education/README.md) on how to add the SDK to your project and authenticate.
