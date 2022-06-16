import com.azure.core.util.Context;
import com.azure.resourcemanager.education.fluent.models.Amount;
import com.azure.resourcemanager.education.fluent.models.StudentDetailsInner;
import com.azure.resourcemanager.education.models.StudentRole;
import java.time.OffsetDateTime;

/** Samples for Students CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/CreateStudent.json
     */
    /**
     * Sample code: Student.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void student(com.azure.resourcemanager.education.EducationManager manager) {
        manager
            .students()
            .createOrUpdateWithResponse(
                "{billingAccountName}",
                "{billingProfileName}",
                "{invoiceSectionName}",
                "{studentAlias}",
                new StudentDetailsInner()
                    .withFirstName("test")
                    .withLastName("user")
                    .withEmail("test@contoso.com")
                    .withRole(StudentRole.STUDENT)
                    .withBudget(new Amount().withCurrency("USD").withValue(100.0f))
                    .withExpirationDate(OffsetDateTime.parse("2021-11-09T22:13:21.795Z")),
                Context.NONE);
    }
}
