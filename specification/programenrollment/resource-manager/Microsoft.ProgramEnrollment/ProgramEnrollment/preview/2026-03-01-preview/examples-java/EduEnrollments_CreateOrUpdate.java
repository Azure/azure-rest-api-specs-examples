
import com.azure.resourcemanager.programenrollment.models.DomainGroup;
import com.azure.resourcemanager.programenrollment.models.EduEnrollmentProperties;
import java.util.Arrays;

/**
 * Samples for EduEnrollments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/EduEnrollments_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update an edu enrollment.
     * 
     * @param manager Entry point to ProgramEnrollmentManager.
     */
    public static void
        createOrUpdateAnEduEnrollment(com.azure.resourcemanager.programenrollment.ProgramEnrollmentManager manager) {
        manager.eduEnrollments().define("default").withRegion("eastus").withExistingResourceGroup("testrg")
            .withProperties(new EduEnrollmentProperties().withDomains(
                Arrays.asList(new DomainGroup().withDomainNames(Arrays.asList("university.edu", "college.edu"))
                    .withTenantId("00000000-0000-0000-0000-000000000001"))))
            .create();
    }
}
