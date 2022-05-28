Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-education_1.0.0-beta.1/sdk/education/azure-resourcemanager-education/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.education.models.InviteCodeGenerateRequest;

/** Samples for Labs GenerateInviteCode. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/GenerateInviteCode.json
     */
    /**
     * Sample code: CreateLab.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void createLab(com.azure.resourcemanager.education.EducationManager manager) {
        manager
            .labs()
            .generateInviteCodeWithResponse(
                "{billingAccountName}",
                "{billingProfileName}",
                "{invoiceSectionName}",
                new InviteCodeGenerateRequest().withMaxStudentCount(10.0f),
                null,
                Context.NONE);
    }
}
```
